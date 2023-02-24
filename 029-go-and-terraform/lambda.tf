// create the lambda function from zip file
resource "aws_lambda_function" "hello_world" {
  function_name = "hello-world"
  description   = "My first hello world function"
  role          = aws_iam_role.lambda.arn
  handler       = local.binary_name
  memory_size   = 128

  filename         = local.archive_path
  source_code_hash = data.archive_file.function_archive.output_base64sha256

  runtime = "go1.x"
}