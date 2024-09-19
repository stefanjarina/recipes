using System;
using System.Collections.Generic;
using System.Text.Json.Serialization;

namespace Recipes.API.entities;

public partial class Measurement
{
    [JsonPropertyName("id")]
    public Guid MeasurementId { get; set; }

    public string Name { get; set; } = null!;

    public Guid TypeId { get; set; }

    public string System { get; set; } = null!;

    public virtual ICollection<Ingredient> Ingredients { get; set; } = new List<Ingredient>();

    public virtual ICollection<StepIngredient> StepIngredients { get; set; } = new List<StepIngredient>();

    public virtual MeasurementType Type { get; set; } = null!;
}
