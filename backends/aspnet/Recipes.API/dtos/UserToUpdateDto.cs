using System.ComponentModel.DataAnnotations;

namespace Recipes.API.dtos;

public class UserToUpdateDto
{
    public string? Email { get; set; }
    
    public string? FullName { get; set; }
    
    public string? Password { get; set; }
}