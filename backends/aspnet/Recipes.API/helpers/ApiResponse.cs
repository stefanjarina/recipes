namespace Recipes.API.helpers;

public static class ApiResponse
{
    public static IResult Ok()
    {
        return Results.Ok();
    }
    
    public static IResult Ok<TValue>(TValue value)
    {
        return Results.Ok(value);
    }
    
    public static IResult BadRequest(string? message = null)
    {
        return Error(message, StatusCodes.Status400BadRequest);
    }
    
    public static IResult Unauthorized(string? message = null)
    {
        return Error(message, StatusCodes.Status401Unauthorized);
    }
    
    public static IResult Forbidden(string? message = null)
    {
        return Error(message, StatusCodes.Status403Forbidden);
    }
    
    public static IResult NotFound(string? message = null)
    {
        return Error(message, StatusCodes.Status404NotFound);
    }
    
    public static IResult Error(string? message = null, int code = StatusCodes.Status500InternalServerError)
    {
        message ??= GetDefaultMessageForStatusCode(code);
        return Results.Json(new { error = message }, statusCode: code);
    }
    
    private static string GetDefaultMessageForStatusCode(int statusCode)
    {
        return statusCode switch
        {
            400 => "A bad request was received.",
            401 => "You are not authorized to perform this action.",
            403 => "You are forbidden from performing this action.",
            404 => "The specified resource could not be found.",
            500 => "An unhandled error occurred. Please try again later.",
            _ => null
        } ?? string.Empty;
    }
}