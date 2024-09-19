using System.Security.Claims;
using AutoMapper;
using Microsoft.EntityFrameworkCore;
using Recipes.API.db;
using Recipes.API.dtos;
using Recipes.API.entities;
using Recipes.API.helpers;
using Recipes.API.extensions;

namespace Recipes.API.handlers;

public static class RecipeHandlers
{
    public static async Task<IResult> GetAllRecipes(RecipesDbContext db)
    {
        var recipes = await db.Recipes.ToListAsync();
        
        return ApiResponse.Ok(recipes);
    }
    
    public static async Task<IResult> GetRecipeById(RecipesDbContext db, Guid id)
    {
        var recipe = await db.Recipes.FindAsync(id);
        
        return recipe == null ? ApiResponse.NotFound() : ApiResponse.Ok(recipe);
    }
    
    public static async Task<IResult> CreateRecipe(RecipesDbContext db, IMapper mapper, IHttpContextAccessor context, RecipeToAddOrUpdateDto recipeDto)
    {
        var idString = context.HttpContext?.User.FindFirst(ClaimTypes.NameIdentifier)?.Value;
        
        if (idString == null)
            return ApiResponse.Unauthorized();
        
        var newRecipe = mapper.Map<Recipe>(recipeDto);
        
        newRecipe.UserId = Guid.Parse(idString);
        
        db.Recipes.Add(newRecipe);
        await db.SaveChangesAsync();
        
        return ApiResponse.Ok(newRecipe);
    }
    
    public static async Task<IResult> UpdateRecipe(RecipesDbContext db, IMapper mapper, IHttpContextAccessor context, Guid id, RecipeToAddOrUpdateDto recipeDto)
    {
        var userId = context.HttpContext?.User.RetrieveIdFromPrincipal();
        
        if (userId == null)
            return ApiResponse.Unauthorized();
        
        var recipe = await db.Recipes.FindAsync(id);
        
        if (recipe == null)
            return ApiResponse.NotFound();
        
        if (recipe.UserId != userId)
            return ApiResponse.Unauthorized();
        
        mapper.Map(recipeDto, recipe);
        
        db.Recipes.Update(recipe);
        await db.SaveChangesAsync();
        
        return ApiResponse.Ok(recipe);
    }
    
    public static async Task<IResult> DeleteRecipe(RecipesDbContext db, IHttpContextAccessor context, Guid id)
    {
        var userId = context.HttpContext?.User.RetrieveIdFromPrincipal();
        
        if (userId == null)
            return ApiResponse.Unauthorized();
        
        var recipe = await db.Recipes.FindAsync(id);
        
        if (recipe == null)
            return ApiResponse.NotFound();
        
        if (recipe.UserId != userId)
            return ApiResponse.Unauthorized();
        
        db.Recipes.Remove(recipe);
        await db.SaveChangesAsync();
        
        return ApiResponse.Ok();
    }
}