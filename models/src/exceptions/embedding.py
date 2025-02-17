class ModelRemoteImportError(Exception):
    def __init__(
        self,
        model_name: str,
        repository: str,
    ) -> None:
        """Exception raised when model failed to import from remote repository

        :param model_name: name of model that failed to import
        :param repository: name of repository used to import model
        """
        self.message = "{name} not found on {repo}"
        self.message = self.message.format(name=model_name, repo=repository)
        super().__init__(self.message)
