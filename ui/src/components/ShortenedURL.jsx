import React from "react";
import { Alert } from "react-bootstrap";
import { Link } from "react-router";
import { BASE_URL } from "../constants/constant";


const ShortenedURL = ({ shortURL }) => {
    if (!shortURL) return null; // Don't render if URL is missing

    return (
        <div className="my-4">
            <Alert key={"light"} variant={"light"}>
                <strong>Shortened URL :</strong> {shortURL}
                <p className="my-1">
                    Click{" "}
                    <a href={`${BASE_URL}/${shortURL}`} target="_blank" rel="noopener noreferrer">
                        here
                    </a>{" "}
                    to redirect
                </p>

                <p>Click <Link to={`/analytics/${shortURL}`}>
                    here
                </Link> for analytics</p>
            </Alert>
        </div>
    );
};

export default ShortenedURL;
