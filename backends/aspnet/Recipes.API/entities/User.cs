using System;
using System.Collections.Generic;
using System.Text.Json.Serialization;

namespace Recipes.API.entities;

public partial class User
{
    [JsonPropertyName("id")]
    public Guid UserId { get; set; }

    public string Email { get; set; } = null!;

    public string FullName { get; set; } = null!;

    [JsonIgnore]
    public string PasswordHash { get; set; } = null!;

    public DateTime? CreatedAt { get; set; }

    public DateTime? UpdatedAt { get; set; }

    [JsonIgnore]
    public virtual ICollection<Comment> Comments { get; set; } = new List<Comment>();

    [JsonIgnore]
    public virtual ICollection<Rating> Ratings { get; set; } = new List<Rating>();

    [JsonIgnore]
    public virtual ICollection<Recipe> Recipes { get; set; } = new List<Recipe>();
}
