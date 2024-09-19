using System;
using System.Collections.Generic;
using System.Text.Json.Serialization;

namespace Recipes.API.entities;

public partial class Ingredient
{
    [JsonPropertyName("id")]
    public Guid IngredientId { get; set; }

    public string Name { get; set; } = null!;

    public decimal Amount { get; set; }

    public decimal AmountMetric { get; set; }

    public Guid? MeasurementId { get; set; }

    public Guid? RecipeId { get; set; }

    public virtual Measurement? Measurement { get; set; }

    public virtual Recipe? Recipe { get; set; }

    public virtual ICollection<StepIngredient> StepIngredients { get; set; } = new List<StepIngredient>();
}
