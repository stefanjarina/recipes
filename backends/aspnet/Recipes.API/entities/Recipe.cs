using System;
using System.Collections.Generic;
using System.Text.Json.Serialization;

namespace Recipes.API.entities;

public partial class Recipe
{
    [JsonPropertyName("id")]
    public Guid RecipeId { get; set; }

    public string Title { get; set; } = null!;

    public string? Description { get; set; }

    public string? Visibility { get; set; }

    public Guid? UserId { get; set; }

    public DateTime? CreatedAt { get; set; }

    public DateTime? UpdatedAt { get; set; }

    public virtual ICollection<Comment> Comments { get; set; } = new List<Comment>();

    public virtual ICollection<Ingredient> Ingredients { get; set; } = new List<Ingredient>();

    public virtual ICollection<Photo> Photos { get; set; } = new List<Photo>();

    public virtual ICollection<Rating> Ratings { get; set; } = new List<Rating>();

    public virtual ICollection<Step> Steps { get; set; } = new List<Step>();

    public virtual User? User { get; set; }
}
