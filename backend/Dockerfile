# Use a minimal distroless image for security and performance
FROM gcr.io/distroless/base-debian11

# Set working directory inside the container
WORKDIR /root/

# Copy the pre-built binary from the host machine
COPY bin/itc .

# Expose the application port (update if needed)
EXPOSE 8080

# Command to run the application
CMD ["./itc"]

