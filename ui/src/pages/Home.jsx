import React, { useState } from "react";
import { Button, Container, Form, InputGroup } from "react-bootstrap";
import ChevronRightIcon from "@mui/icons-material/ChevronRight";
import axios from "axios";
import { useMutation } from "@tanstack/react-query";

import { Bounce, ToastContainer, toast } from 'react-toastify';
import ShortenedURL from "../components/ShortenedURL";

import Spinner from 'react-bootstrap/Spinner';



const Home = () => {
    const [url, setURL] = useState("");
    const [alias, setAlias] = useState("");
    const [ttl, setTTL] = useState(120);
    const [shortURL, setShortURL] = useState("");

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
        },
        onError: (error) => {
            toast.error('Error shortening URL')
            console.error("Error shortening URL:", error);
        },
    });

    const submitHandler = (e) => {
        e.preventDefault();
        mutation.mutate();
    };

    return (
        <Container className="url-cont">
            <h1 className="my-5">URL Shortener</h1>
            <Form onSubmit={submitHandler}>
                <InputGroup className="my-2">
                    <Form.Control type="text" placeholder="Enter the link here" value={url} onChange={(e) => setURL(e.target.value)} required />
                </InputGroup>

                <InputGroup className="my-2">
                    <div className="d-flex gap-3 w-100">
                        <Form.Control type="text" placeholder="Enter the custom alias (optional)" value={alias} onChange={(e) => setAlias(e.target.value)} />
                        <Form.Control type="number" placeholder="Enter the TTL in seconds" value={ttl} onChange={(e) => setTTL(Number(e.target.value))} />
                    </div>
                </InputGroup>

                <Button variant="primary" type="submit" className="my-4" disabled={mutation.isLoading}>
                    {mutation.isLoading ? "Shortening..." : "Shorten"} <ChevronRightIcon />
                </Button>
            </Form>

            {mutation.status === "pending" && <>
                <div className="d-flex justify-content-center my-3">
                    <Spinner animation="border" role="status">
                        <span className="visually-hidden">Loading...</span>
                    </Spinner>
                </div>
            </>}


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
