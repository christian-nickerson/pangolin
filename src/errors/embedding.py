class ModelRemoteImportError(Exception):

    def __init__(
        self,
        model_name: str,
        repository: str,
    ) -> None:
        """Exception raised when model failed to import from remote repository

        :param model_name: name of model that failed to import
        :param site: name of repository used to import model
        :param message: explanation of the error, defaults to "{model_name} could not be found"
        """
        self.message = "{model_name} not found on {repository}"
        self.message = self.message.format(model_name, repository)
        super().__init__(self.message)
