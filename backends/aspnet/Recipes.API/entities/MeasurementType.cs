using System;
using System.Collections.Generic;
using System.Text.Json.Serialization;

namespace Recipes.API.entities;

public partial class MeasurementType
{
    [JsonPropertyName("id")]
    public Guid TypeId { get; set; }

    public string Name { get; set; } = null!;

    public virtual ICollection<Measurement> Measurements { get; set; } = new List<Measurement>();
}
