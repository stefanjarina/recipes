﻿using System.ComponentModel.DataAnnotations;

namespace Recipes.API.dtos;

public class LoginDto
{
    [Required]
    public string Email { get; set; } = null!;

    [Required]
    public string Password { get; set; } = null!;
}