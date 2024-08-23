
# TerranSys Atlas

**TerranSys Atlas** is the module responsible for creating and managing the foundational infrastructure on AWS required to deploy an EKS cluster. Using Pulumi and Go, this module automates the setup of VPCs, subnets, IAM roles, EKS clusters, and more.

## Prerequisites

Before working with **TerranSys Atlas**, make sure you have the following tools installed:

- **Go 1.18+**: [Install Go](https://golang.org/doc/install)
- **Pulumi CLI**: [Install Pulumi](https://www.pulumi.com/docs/get-started/install/)
- **Git**: [Install Git](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git)
- **AWS CLI**: [Install AWS CLI](https://docs.aws.amazon.com/cli/latest/userguide/install-cliv2.html)
- **GitHub CLI (gh)**: [Install GitHub CLI](https://cli.github.com/)

Additionally, ensure your AWS credentials are configured:

```bash
aws configure
```

## Installation

Follow these steps to clone the repository and install the necessary dependencies:

1. **Clone the repository:**

   ```bash
   git clone https://github.com/terransys/terransys-atlas.git
   cd terransys-atlas
   ```

2. **Install Go dependencies:**

   ```bash
   go mod init terransys-atlas
   go mod tidy
   ```

3. **Install Pulumi dependencies:**

   If this is your first time using Pulumi in this project, Pulumi will automatically download the necessary plugins the first time you run `pulumi up`. However, if you want to install them manually:

   ```bash
   pulumi plugin install resource aws v5.0.0
   ```

## Configuration

### Pulumi Configuration

Pulumi uses stacks to manage environment-specific configurations (e.g., `dev`, `prod`).

1. **Select the development stack:**

   ```bash
   pulumi stack select dev
   ```

2. **Set environment variables in Pulumi:**

   Configure necessary variables, such as the AWS region:

   ```bash
   cd  cmd/atlas
   pulumi config set aws:region us-west-2 --stack dev
   ```

3. **Add other necessary configurations:**

   For example, if your project requires additional specific settings:

   ```bash
   pulumi config set eksNodeInstanceType t3.medium --stack dev
   ```

## Deploying the Project

To deploy the infrastructure:

1. **Run Pulumi:**

   ```bash
   pulumi up
   ```

   Review the changes to be applied and confirm the deployment. Pulumi will handle creating and configuring all necessary resources in AWS.

2. **Check the State:**

   To view the state of the deployed resources:

   ```bash
   pulumi stack output
   ```

3. **Destroy the Infrastructure (optional):**

   If you need to destroy the created infrastructure:

   ```bash
   pulumi destroy
   ```

## Project Structure

This project is organized into Go packages, each handling different aspects of the infrastructure:

```
cmd/
├── atlas/                    # Entry point for running Pulumi
internal/
├── eksroles/                 # Module for configuring IAM roles and policies
├── eksnet/                   # Module for configuring EKS networking
├── ekscluster/               # Module for configuring the EKS cluster
└── eksinit/                  # Module for initial cluster configurations
pkg/                          # Exportable, reusable packages
```

## Contributing

We'd love for you to contribute to **TerranSys**! Follow these steps to get started:

1. **Fork the repository:**
   - Fork the repository on GitHub.

2. **Clone your fork locally:**

   ```bash
   git clone https://github.com/your-username/terransys-atlas.git
   cd terransys-atlas
   ```

3. **Create a new branch for your feature or fix:**

   ```bash
   git checkout -b my-new-feature
   ```

4. **Make your changes and commit them:**

   Make sure to follow Go's style conventions and document your changes.

   ```bash
   git add .
   git commit -m "Add new feature X"
   ```

5. **Push your branch:**

   ```bash
   git push origin my-new-feature
   ```

6. **Create a Pull Request:**
   - Go to your fork's GitHub page and open a Pull Request to the original repository.

## Best Practices

- **Write Tests:** Ensure any new functionality is well-tested.
- **Documentation:** Update documentation when adding new features.
- **Linting and Formatting:** Use `gofmt` and `golint` to keep the code clean and consistent.

## License

**TerranSys Atlas** is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
