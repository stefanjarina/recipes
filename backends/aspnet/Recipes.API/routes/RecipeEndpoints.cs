using Recipes.API.handlers;

namespace Recipes.API.routes;

public static class RecipeEndpoints
{
    public static RouteGroupBuilder MapRecipeRoutes(this RouteGroupBuilder group)
    {
        group.MapGet("/recipes", RecipeHandlers.GetAllRecipes)
            .WithOpenApi();

        group.MapGet("/recipes/{id:guid}", RecipeHandlers.GetRecipeById)
            .WithOpenApi();

        group.MapPost("/recipes", RecipeHandlers.CreateRecipe)
            .RequireAuthorization()
            .WithOpenApi();
        
        group.MapPatch("/recipes/{id:guid}", RecipeHandlers.UpdateRecipe)
            .RequireAuthorization()
            .WithOpenApi();
        
        group.MapDelete("/recipes/{id:guid}", RecipeHandlers.DeleteRecipe)
            .RequireAuthorization()
            .WithOpenApi();

        return group;
    }
}