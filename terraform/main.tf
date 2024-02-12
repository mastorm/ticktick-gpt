terraform {
  required_providers {
    aws = {
      source = "hashicorp/aws"
    }
    archive = {
      source = "hashicorp/archive"
    }
    null = {
      source = "hashicorp/null"
    }
  }

  required_version = ">= 1.3.7"
}


provider "aws" {
  region = "eu-west-1"

  default_tags {
    tags = {
      app = "ticktick-gpt"
    }
  }
}

data "aws_iam_policy_document" "assume_lambda_role" {
  statement {
    actions = ["sts:AssumeRole"]

    principals {
      type        = "Service"
      identifiers = ["lambda.amazonaws.com"]
    }
  }
}

resource "aws_iam_role" "lambda" {
  name               = "AssumeLambdaRole"
  description        = "Role for lambda to assume lambda"
  assume_role_policy = data.aws_iam_policy_document.assume_lambda_role.json
}

data "aws_iam_policy_document" "allow_lambda_logging" {
  statement {
    effect = "Allow"
    actions = [
      "logs:CreateLogStream",
      "logs:PutLogEvents",
    ]

    resources = [
      "arn:aws:logs:*:*:*",
    ]
  }
}

resource "aws_iam_policy" "function_logging_policy" {
  name        = "AllowLambdaLoggingPolicy"
  description = "Policy for lambda cloudwatch logging"
  policy      = data.aws_iam_policy_document.allow_lambda_logging.json
}

resource "aws_iam_role_policy_attachment" "lambda_logging_policy_attachment" {
  role       = aws_iam_role.lambda.id
  policy_arn = aws_iam_policy.function_logging_policy.arn
}

resource "null_resource" "function_binary" {
  provisioner "local-exec" {
    command = "cross-env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 GOFLAGS=-trimpath go build -mod=readonly -o ${local.binary_path} ${local.src_path}"
  }
}

data "archive_file" "function_archive" {
  depends_on = [null_resource.function_binary]

  type        = "zip"
  source_file = local.binary_path
  output_path = local.archive_path
}

resource "aws_lambda_function" "function" {
  function_name = local.function_name
  description   = ""
  role          = aws_iam_role.lambda.arn
  handler       = local.binary_name
  memory_size   = 128



  filename         = local.archive_path
  source_code_hash = data.archive_file.function_archive.output_base64sha256

  runtime = "go1.x"
}

resource "aws_lambda_function_url" "test_latest" {
  function_name      = aws_lambda_function.function.function_name
  authorization_type = "NONE"
}


resource "aws_cloudwatch_log_group" "log_group" {
  name              = "/aws/lambda/${aws_lambda_function.function.function_name}"
  retention_in_days = 1
}
