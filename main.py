import toml

with open("config.toml", "r") as f:
    config = toml.load(f)

print("Title:", config["title"])
print("Author:", config["author"]["name"])
print("Email:", config["author"]["email"])
