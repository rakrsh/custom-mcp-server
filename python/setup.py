from setuptools import setup, find_packages

setup(
    name="custom_mcp_server",
    version="0.1.0",
    description="Python implementation of the Custom AI Agents MCP server",
    packages=find_packages(include=["custom_mcp_server", "custom_mcp_server.*"]),
    python_requires=">=3.11",
    classifiers=[
        "Programming Language :: Python :: 3",
        "Programming Language :: Python :: 3.11",
        "Programming Language :: Python :: 3.12",
        "License :: OSI Approved :: MIT License",
        "Operating System :: OS Independent",
    ],
)
