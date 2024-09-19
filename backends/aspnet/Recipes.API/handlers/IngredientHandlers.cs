using Microsoft.EntityFrameworkCore;
using Recipes.API.db;
using Recipes.API.helpers;

namespace Recipes.API.handlers;

public static class IngredientHandlers
{
    public static async Task<IResult> GetIngredientsAsync(RecipesDbContext db, Guid id)
    {
        var recipe = await db.Recipes
            .Where(r => r.RecipeId == id)
            .Include(r => r.Ingredients)
            .FirstOrDefaultAsync();
        
        if (recipe == null)
            return ApiResponse.NotFound();

        return ApiResponse.Ok(recipe.Ingredients);
    }
}