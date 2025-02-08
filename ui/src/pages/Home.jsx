import React, { useState } from "react";
import { Button, Container, Form, InputGroup } from "react-bootstrap";
import ChevronRightIcon from "@mui/icons-material/ChevronRight";
import axios from "axios";
import { useMutation } from "@tanstack/react-query";
import { Bounce, ToastContainer, toast } from 'react-toastify';
import ShortenedURL from "../components/ShortenedURL";
import Spinner from 'react-bootstrap/Spinner';
import Error from "../components/Error";

const Home = () => {
    const [url, setURL] = useState("");
    const [alias, setAlias] = useState("");
    const [ttl, setTTL] = useState(120);
    const [shortURL, setShortURL] = useState("");
    const [error, setError] = useState(null);

    const mutation = useMutation({
        mutationFn: async () => {
            const response = await axios.post("http://localhost:8080/shorten", {
                long_url: url,
                custom_alias: alias || "",
                ttl_seconds: ttl,
            });
            return response.data;
        },
        onSuccess: (data) => {
            setShortURL(data.short_url || "");
            setError(null);
        },
        onError: (err) => {
            const errorMessage = err.response?.data || 'Error shortening URL';
            // const errorStatus = err.response?.status || 'Unknown Status';
            // const errorDetails = err.response?.statusText || 'No Details';

            console.error("Full error response:", err.response);
            // toast.error(`Error: ${errorMessage} (Status: ${errorStatus} - ${errorDetails})`);
            setError(errorMessage);
        },

    });

    const submitHandler = (e) => {
        e.preventDefault();
        setError(null);
        mutation.mutate();
    };

    return (
        <Container className="url-cont">
            <h1 className="my-5">URL Shortener</h1>
            <Form onSubmit={submitHandler}>
                <InputGroup className="my-2">
                    <Form.Control
                        type="text"
                        placeholder="Enter the link here"
                        value={url}
                        onChange={(e) => setURL(e.target.value)}
                        required
                    />
                </InputGroup>

                <InputGroup className="my-2">
                    <div className="d-flex gap-3 w-100">
                        <Form.Control
                            type="text"
                            placeholder="Enter the custom alias (optional)"
                            value={alias}
                            onChange={(e) => setAlias(e.target.value)}
                        />
                        <Form.Control
                            type="number"
                            placeholder="Enter the TTL in seconds"
                            value={ttl}
                            onChange={(e) => setTTL(Number(e.target.value))}
                        />
                    </div>
                </InputGroup>

                <Button variant="primary" type="submit" className="my-4" disabled={mutation.isLoading}>
                    {mutation.isLoading ? "Shortening..." : "Shorten"} <ChevronRightIcon />
                </Button>
            </Form>

            {mutation.status === "pending" && (
                <div className="d-flex justify-content-center my-3">
                    <Spinner animation="border" role="status">
                        <span className="visually-hidden">Loading...</span>
                    </Spinner>
                </div>
            )}

            {mutation.isError && error && <Error error={error} />}

            {!mutation.isError && mutation.data && <ShortenedURL shortURL={shortURL} />}

            <ToastContainer
                position="bottom-center"
                autoClose={5000}
                hideProgressBar={false}
                newestOnTop={false}
                closeOnClick
                rtl={false}
                pauseOnFocusLoss
                draggable
                pauseOnHover
                theme="light"
                transition={Bounce}
            />
        </Container>
    );
};

export default Home;
