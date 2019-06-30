import pandas as pd
import os
import sys
import re

# def get_allergy_info(item, recipe_no):
#     allergy_info = []
#     allergy = item.split('.')
#     for ind, recipe in enumerate(recipes):
#         index = recipe.find('allergy')
#         if(index == -1):
#             continue
#         else:
#             print(recipe)
#             sentances = recipe.split('.')
#             print(sentances,end="\n")
#             allergy = [sentance for sentance in sentances if "allergy" in sentance]
#             allergy_info.append([ind, allergy[0]])
#             recipes[ind] = '.'.join(sentances.remove(sentances.index(allergy[0])))
#     return allergy_info, recipes


def remove_brack(item):
    return item.split('(')[0]


def remove_dot(item):
    return item.split('.')[0]


def clean_item(item):
    return remove_brack(remove_dot(item)).split('and')

def get_items_from_recipe(recipe):
    items = [clean_item(string)
                for string in recipe.split(',')]
    # items.sort()
    return items


def get_items(filepath,savefolder):
    all_items=[]
    data = pd.read_csv(filepath).dropna(subset=["features.value", "name"])
    print(data.head())
    recipes = data["features.value"].str.lower().tolist()
    # allergy_info, recipes = get_allergy_info(recipes)
    recipe_table=[]
    for recipe_no,recipe in enumerate(recipes):
        recipe = re.sub("[\(\[].*?[\)\]]", "", recipe)
        recipe_items=[]
        for items in get_items_from_recipe(recipe):
            for sep_items in items:
                all_items.append(sep_items)
                recipe_items.append(sep_items)
        recipe_table.append((recipe_items,recipe_no))
    for item,recipe in zip(recipes[:10],recipe_table[:10]):
        print(item,recipe,sep="\n")
        print('-'*90)
    pd.DataFrame(recipe_table,columns=["recipe","recipe_number"]).to_csv(savefolder+"/recipe_table.csv",index=False)
    pd.DataFrame({'items':all_items}).to_csv(savefolder+"/items.csv")


if __name__ == "__main__":
    if(len(sys.argv) != 3):
        print("Give File paths")
        quit()
    filepath = sys.argv[1]
    savefolder = sys.argv[2]
    print(filepath, savefolder)
    if(os.path.exists(filepath) and os.path.exists(savefolder)):
        get_items(filepath,savefolder)
    else:
        print("Path does not exist")
        quit()
