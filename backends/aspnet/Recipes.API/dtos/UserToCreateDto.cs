using System.ComponentModel.DataAnnotations;

namespace Recipes.API.dtos;

public class UserToCreateDto
{
    [Required]
    public string Email { get; set; } = null!;
    
    [Required]
    public string FullName { get; set; } = null!;
    
    [Required]
    public string Password { get; set; } = null!;
}