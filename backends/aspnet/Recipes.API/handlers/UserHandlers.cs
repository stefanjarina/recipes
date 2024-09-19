using Microsoft.EntityFrameworkCore;
using Recipes.API.db;

namespace Recipes.API.handlers;

public static class UserHandlers
{
    public static async Task<IResult> GetAllUsers(RecipesDbContext db)
    {
        var users = await db.Users.ToListAsync();
        
        return Results.Ok(users);
    }
    
    public static async Task<IResult> GetUserById(RecipesDbContext db, Guid id)
    {
        var user = await db.Users.FindAsync(id);
        
        if (user == null)
        {
            return Results.NotFound();
        }
        
        return Results.Ok(user);
    }
    
    public static async Task<IResult> GetUserByEmail(RecipesDbContext db, string email)
    {
        var user = await db.Users.FirstOrDefaultAsync(u => u.Email == email);
        
        if (user == null)
        {
            return Results.NotFound();
        }
        
        return Results.Ok(user);
    }
}