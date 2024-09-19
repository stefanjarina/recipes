using System;
using System.Collections.Generic;
using System.Text.Json.Serialization;

namespace Recipes.API.entities;

public partial class Photo
{
    [JsonPropertyName("id")]
    public Guid PhotoId { get; set; }

    public Guid? RecipeId { get; set; }

    public string? FilePath { get; set; }

    public string? Url { get; set; }

    public DateTime? CreatedAt { get; set; }

    public DateTime? UpdatedAt { get; set; }

    public virtual Recipe? Recipe { get; set; }
}
