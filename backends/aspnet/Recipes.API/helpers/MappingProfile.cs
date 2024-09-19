using AutoMapper;
using Recipes.API.dtos;
using Recipes.API.entities;

namespace Recipes.API.helpers;

public class MappingProfile : Profile
{
    public MappingProfile()
    {
        CreateMap<UserToCreateDto, User>();
        CreateMap<RecipeToAddOrUpdateDto, Recipe>()
            .ForAllMembers(opts => 
            {
                opts.AllowNull();
                opts.Condition((src, dest, srcMember) => srcMember != null);
            });
    }
}