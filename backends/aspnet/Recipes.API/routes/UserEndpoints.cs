using Recipes.API.handlers;

namespace Recipes.API.routes;

public static class UserEndpoints
{
    public static RouteGroupBuilder MapUserRoutes(this RouteGroupBuilder group)
    {
        group.MapGet("/users", UserHandlers.GetAllUsers)
            .RequireAuthorization()
            .WithOpenApi();

        group.MapGet("/users/{id:guid}", UserHandlers.GetUserById)
            .WithOpenApi();

        group.MapPost("/users", UserHandlers.CreateUser)
            .WithOpenApi();

        group.MapPatch("/users/{id:guid}", UserHandlers.UpdateUser)
            .RequireAuthorization()
            .WithOpenApi();

        group.MapDelete("/users/{id:guid}", UserHandlers.DeleteUser)
            .RequireAuthorization()
            .WithOpenApi();

        return group;
    }
}