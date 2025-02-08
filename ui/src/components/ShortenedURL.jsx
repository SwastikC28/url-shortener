import React from "react";
import { Alert } from "react-bootstrap";

const ShortenedURL = ({ shortURL }) => {
    if (!shortURL) return null; // Don't render if URL is missing

    return (
        <div className="my-4">
            <Alert key={"light"} variant={"light"}>
                <strong>Shortened URL:</strong> {shortURL}
                <p className="my-1">
                    Click{" "}
                    <a href={shortURL} target="_blank" rel="noopener noreferrer">
                        here
                    </a>{" "}
                    to redirect
                </p>
            </Alert>
        </div>
    );
};

export default ShortenedURL;
