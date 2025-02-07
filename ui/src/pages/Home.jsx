import React, { useState } from 'react';
import { Button, Container, Form, InputGroup } from 'react-bootstrap';
import ChevronRightIcon from '@mui/icons-material/ChevronRight';

const Home = () => {

    const [url, setURL] = useState('')

    const submitHandler = (e) => {
        e.preventDefault()

        console.log(url);
        
    }

    const onChangeHandler = (e) => {
        setURL(e.target.value)
    }

    return (
        <Container className='url-cont'>
            <h1 className='my-5'>URL Shortener</h1>
            <Form onSubmit={submitHandler}>
                <InputGroup>
                    <Form.Control type="text" placeholder="Enter the link here" value={url} onChange={onChangeHandler} />
                    <Button variant="primary" type="submit">
                        <ChevronRightIcon />
                    </Button>
                </InputGroup>
            </Form>
        </Container>
    );
}

export default Home;
