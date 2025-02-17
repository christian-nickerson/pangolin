from dynaconf import Dynaconf

settings = Dynaconf(
    envvar_prefix="PANGOLIN",
    settings_files=["settings.toml", ".secrets.toml"],
)
