variable "myAge" {
    description = "This is how old I am"
    type = int
    default = 21
    validation {
        condition = myAge >= 23
        error_message = "This is the wrong age"
        }
}