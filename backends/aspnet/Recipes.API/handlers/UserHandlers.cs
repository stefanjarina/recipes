using System.Security.Claims;
using AutoMapper;
using Microsoft.EntityFrameworkCore;
using Recipes.API.db;
using Recipes.API.dtos;
using Recipes.API.entities;
using Recipes.API.extensions;
using Recipes.API.helpers;

namespace Recipes.API.handlers;

public static class UserHandlers
{
    public static async Task<IResult> GetAllUsers(RecipesDbContext db)
    {
        var users = await db.Users.ToListAsync();
        
        return ApiResponse.Ok(users);
    }
    
    public static async Task<IResult> GetUserById(RecipesDbContext db, Guid id)
    {
        var user = await db.Users.FindAsync(id);
        
        return user == null ? ApiResponse.NotFound() : ApiResponse.Ok(user);
    }

    public static async Task<IResult> CreateUser(RecipesDbContext db, IMapper mapper, UserToCreateDto user)
    {
        
        var passwordHash = BCrypt.Net.BCrypt.EnhancedHashPassword(user.Password);
        
        var newUser = mapper.Map<User>(user);

        newUser.PasswordHash = passwordHash;
        
        db.Users.Add(newUser);
        await db.SaveChangesAsync();
        
        return ApiResponse.Ok(newUser);
    }
    
    public static async Task<IResult> UpdateUser(RecipesDbContext db, IHttpContextAccessor context, Guid id, UserToUpdateDto user)
    {
        var userId = context.HttpContext?.User.RetrieveIdFromPrincipal();
        
        if (userId != id)
            return ApiResponse.Unauthorized();
        
        if (userId != id)
            return ApiResponse.Forbidden();
        
        var existingUser = await db.Users.FindAsync(id);
        
        if (existingUser == null)
            return ApiResponse.NotFound();
        
        if (!string.IsNullOrWhiteSpace(user.Email))
            existingUser.Email = user.Email;
        
        if (!string.IsNullOrWhiteSpace(user.FullName))
            existingUser.FullName = user.FullName;
        
        if (!string.IsNullOrWhiteSpace(user.Password))
            existingUser.PasswordHash = BCrypt.Net.BCrypt.EnhancedHashPassword(user.Password);
        
        await db.SaveChangesAsync();
        
        return ApiResponse.Ok(existingUser);
    }
    
    public static async Task<IResult> DeleteUser(RecipesDbContext db, IHttpContextAccessor context, Guid id)
    {
        var idFromToken = context.HttpContext?.User.FindFirst(ClaimTypes.NameIdentifier)?.Value;
        
        if (idFromToken != id.ToString())
            return ApiResponse.Forbidden();
        
        var existingUser = await db.Users.FindAsync(id);
        
        if (existingUser == null)
            return ApiResponse.NotFound();
        
        db.Users.Remove(existingUser);
        await db.SaveChangesAsync();
        
        return ApiResponse.Ok();
    }
}