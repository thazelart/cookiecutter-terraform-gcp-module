/**
 * # {{cookiecutter.module_name}}
 * {{ cookiecutter.short_description }}
 * 
 * ## Documentation
 * In order to generate the documentation, run [terraform-docs](https://github.com/segmentio/terraform-docs)
 * 
 * ```
 * terraform-docs markdown tables . > README.md
 * ```
 * 
 * ## Usage
 * You can either call it with parameters to define what your network must look like :
 * ```hcl
 * module "{{cookiecutter.module_name}}" {
 *   source = "git@github.com:{{cookiecutter.github_user}}/{{cookiecutter.module_name}}"
 *   gcp_project_id = "${var.project_id}"
 * ```
 */
