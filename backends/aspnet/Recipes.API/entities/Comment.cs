using System;
using System.Collections.Generic;
using System.Text.Json.Serialization;

namespace Recipes.API.entities;

public partial class Comment
{
    [JsonPropertyName("id")]
    public Guid CommentId { get; set; }

    public Guid? RecipeId { get; set; }

    public Guid? UserId { get; set; }

    public string CommentText { get; set; } = null!;

    public DateTime? CreatedAt { get; set; }

    public DateTime? UpdatedAt { get; set; }

    public virtual Recipe? Recipe { get; set; }

    public virtual User? User { get; set; }
}
