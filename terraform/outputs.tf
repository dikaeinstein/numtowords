output "bucket_domain_name" {
  description = "The bucket domain name. Will be of format bucketname.s3.amazonaws.com"
  value       = aws_s3_bucket.site.bucket_domain_name
}

output "website_bucket_name" {
  description = "Name (id) of the bucket"
  value       = aws_s3_bucket.site.id
}

output "website_endpoint" {
  description = "Website endpoint"
  value       = "http://${aws_s3_bucket.site.website_endpoint}"
}
