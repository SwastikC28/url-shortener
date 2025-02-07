import React, { use, useState } from 'react';
import { Button, Col, Container, Form, InputGroup } from 'react-bootstrap';
import ChevronRightIcon from '@mui/icons-material/ChevronRight';
import axios from 'axios'

const Home = () => {

    const [url, setURL] = useState('')
    const [alias, setAlias] = useState('')
    const [ttl, setTTL] = useState(120)

    const [shortURL, setShortURL] = useState('')

    const submitHandler = async (e) => {
        e.preventDefault()

        try {
            const response = await axios.post("http://localhost:8080/shorten", {
                "long_urls": url,
                "custom_alias": "",
                "ttl_seconds": 120,
            })

            console.log(response.data);
            
            // const data=response.data
            // if data
        } catch (error) {

        }
    }

    const onURLChangeHandler = (e) => {
        setURL(e.target.value)
    }

    const onAliasChangeHandler = (e) => {
        setAlias(e.target.value)
    }

    const onTTLChangeHandler = (e) => {
        setTTL(e.target.value)
    }

    return (
        <Container className='url-cont'>
            <h1 className='my-5'>URL Shortener</h1>
            <Form onSubmit={submitHandler}>
                <InputGroup className='my-2'>
                    <Form.Control type="text" placeholder="Enter the link here" value={url} onChange={onURLChangeHandler} />
                </InputGroup>

                <InputGroup className="my-2">
                    <div className="d-flex gap-3 w-100">
                        <Form.Control type="text" placeholder="Enter the custom alias" value={alias} onChange={onAliasChangeHandler} />
                        <Form.Control type="number" placeholder="Enter the TTL in seconds" value={ttl} onChange={onTTLChangeHandler} />
                    </div>
                </InputGroup>

                <Button variant="primary" type="submit" className='my-4'>
                    Shorten <ChevronRightIcon />
                </Button>
            </Form>
        </Container>
    );
}

export default Home;
