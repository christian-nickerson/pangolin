from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import (
    ClassVar as _ClassVar,
    Iterable as _Iterable,
    Mapping as _Mapping,
    Optional as _Optional,
    Union as _Union,
)

DESCRIPTOR: _descriptor.FileDescriptor

class InferenceRequest(_message.Message):
    __slots__ = ("text", "model_name")
    TEXT_FIELD_NUMBER: _ClassVar[int]
    MODEL_NAME_FIELD_NUMBER: _ClassVar[int]
    text: _containers.RepeatedScalarFieldContainer[str]
    model_name: str
    def __init__(self, text: _Optional[_Iterable[str]] = ..., model_name: _Optional[str] = ...) -> None: ...

class Vector(_message.Message):
    __slots__ = ("components",)
    COMPONENTS_FIELD_NUMBER: _ClassVar[int]
    components: _containers.RepeatedScalarFieldContainer[float]
    def __init__(self, components: _Optional[_Iterable[float]] = ...) -> None: ...

class InferenceResponse(_message.Message):
    __slots__ = ("embeddings",)
    EMBEDDINGS_FIELD_NUMBER: _ClassVar[int]
    embeddings: _containers.RepeatedCompositeFieldContainer[Vector]
    def __init__(self, embeddings: _Optional[_Iterable[_Union[Vector, _Mapping]]] = ...) -> None: ...

class ModelListRequest(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class ModelListResponse(_message.Message):
    __slots__ = ("model_names",)
    MODEL_NAMES_FIELD_NUMBER: _ClassVar[int]
    model_names: _containers.RepeatedScalarFieldContainer[str]
    def __init__(self, model_names: _Optional[_Iterable[str]] = ...) -> None: ...
