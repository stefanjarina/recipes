using Recipes.API.handlers;

namespace Recipes.API.routes;

public static class AuthEndpoints
{
    public static RouteGroupBuilder MapAuthRoutes(this RouteGroupBuilder group)
    {
        group.MapPost("/login", AuthHandlers.LoginAsync)
            .WithOpenApi();

        return group;
    }
}