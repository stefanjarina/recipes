using System.Text.Json.Serialization;

namespace Recipes.API.helpers;

public class ApiError(string error)
{
    [JsonPropertyName("error")]
    public string Error { get; set; } = error;
}