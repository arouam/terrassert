resource "terrassert_bool" "check_valid" {
  name = "terrassert bool"
  condition = false
  error_message = "for a fargate launch type a network config is required"
}