using System.Security.Claims;

namespace Recipes.API.extensions;

public static class ClaimsPrincipalExtensions
{
    public static string RetrieveEmailFromPrincipal(this ClaimsPrincipal user)
    {
        return user?.Claims?.FirstOrDefault(x => x.Type == ClaimTypes.Email)?.Value!;
    }
    
    public static Guid RetrieveIdFromPrincipal(this ClaimsPrincipal user)
    {
        var idString = user?.FindFirst(ClaimTypes.NameIdentifier)?.Value!;
        return Guid.Parse(idString);
    }
}