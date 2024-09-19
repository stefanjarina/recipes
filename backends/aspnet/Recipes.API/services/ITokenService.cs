using Recipes.API.entities;

namespace Recipes.API.services;

public interface ITokenService
{
    string GenerateToken(User user);
}