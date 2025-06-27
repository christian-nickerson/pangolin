# Pangolin

### What is Pangolin?

Pangolin is a fast and scalable context store that allows applications, LLMs and users to read, write and search text and documents.

Pangolin has an opionanted data model, allowing users to create and configure collections of documents. Each collection is generated with a specification for the emebedding model, chunk size and distance metric that is used to both process and search documents. Pangolin's "batteries included" model allows users simply write documents to the content store, allowing Pangolin to process and index each document.

Pangolin provides both an API and MCP interface to allow classic applications and Agents to interact with the service.

Pangolin uses a distributed model, in both compute and storage, to allow the application to horizontally scale. Pangolin makes use of S3 compatible object stores, like MinIO, to store data and metadata about the application.

Pangolin also provides a TUI client to allow users and administrators to monitor, review and interact with the context store.
