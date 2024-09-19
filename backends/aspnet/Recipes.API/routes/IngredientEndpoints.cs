using Recipes.API.handlers;

namespace Recipes.API.routes;

public static class IngredientEndpoints
{
    public static RouteGroupBuilder MapIngredientRoutes(this RouteGroupBuilder group)
    {
        group.MapGet("/recipes/{id:guid}/ingredients", IngredientHandlers.GetIngredientsAsync)
            .WithOpenApi();

        return group;
    }
}