using System;
using System.Collections.Generic;
using System.Text.Json.Serialization;

namespace Recipes.API.entities;

public partial class Step
{
    [JsonPropertyName("id")]
    public Guid StepId { get; set; }

    public string Title { get; set; } = null!;

    public string Instructions { get; set; } = null!;

    public int StepNumber { get; set; }

    public Guid? RecipeId { get; set; }

    public DateTime? CreatedAt { get; set; }

    public DateTime? UpdatedAt { get; set; }

    public virtual Recipe? Recipe { get; set; }

    public virtual ICollection<StepIngredient> StepIngredients { get; set; } = new List<StepIngredient>();
}
