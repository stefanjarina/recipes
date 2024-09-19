using System.ComponentModel.DataAnnotations;

namespace Recipes.API.dtos;

public class RecipeToAddOrUpdateDto
{
    [Required]
    public string Title { get; set; } = null!;

    public string? Description { get; set; }

    [Required]
    public string Visibility { get; set; } = "public";
}