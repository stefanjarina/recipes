using System;
using System.Collections.Generic;
using System.Text.Json.Serialization;

namespace Recipes.API.entities;

public partial class Rating
{
    [JsonPropertyName("id")]
    public Guid RatingId { get; set; }

    public Guid? UserId { get; set; }

    public Guid? RecipeId { get; set; }

    public int? Rating1 { get; set; }

    public DateTime? CreatedAt { get; set; }

    public DateTime? UpdatedAt { get; set; }

    public virtual Recipe? Recipe { get; set; }

    public virtual User? User { get; set; }
}
