using System;
using System.Collections.Generic;
using System.Text.Json.Serialization;

namespace Recipes.API.entities;

public partial class StepIngredient
{
    [JsonPropertyName("id")]
    public Guid StepIngredientId { get; set; }

    public Guid? StepId { get; set; }

    public Guid? IngredientId { get; set; }

    public decimal Amount { get; set; }

    public decimal AmountMetric { get; set; }

    public Guid? MeasurementId { get; set; }

    public virtual Ingredient? Ingredient { get; set; }

    public virtual Measurement? Measurement { get; set; }

    public virtual Step? Step { get; set; }
}
