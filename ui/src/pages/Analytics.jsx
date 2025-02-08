import { useQuery } from '@tanstack/react-query'
import React from 'react'
import { Button, Col, Container, Row, Spinner } from 'react-bootstrap'
import { Link, useParams } from 'react-router'
import Timestamps from '../components/Timestamps'
import { BASE_URL } from '../constants/constant'

const Analytics = () => {
    const { shortURL } = useParams()

    const { isPending, error, data, isFetching, isLoading } = useQuery({
        queryKey: ['analyticsData', shortURL], // <-- Make query key dynamic
        queryFn: async () => {
            const response = await fetch(
                `${BASE_URL}/analytics/${shortURL}`
            )
            return await response.json()
        },
        enabled: !!shortURL,
    });

    return (
        <Container className="url-cont analytics-cont">
            <Row className='mb-4'>
                <Col>
                    <Button as={Link} to={'/'} variant='secondary'>Home</Button>
                </Col>
            </Row>

            <Row className='mb-2'>
                <h1>Analytics</h1>
            </Row>

            {isLoading && <div className="d-flex justify-content-center my-5">
                <Spinner animation="border" role="status">
                    <span className="visually-hidden">Loading...</span>
                </Spinner>
            </div>}

            {data && <>
                <Row className='my-1'>
                    <Col>
                        Short URL
                    </Col>

                    <Col>
                        {data.short_url}
                    </Col>
                </Row>

                <Row className='my-1'>
                    <Col>
                        Long URL
                    </Col>

                    <Col>
                        {data.long_url}
                    </Col>
                </Row>

                <Row className='my-1'>
                    <Col>
                        Access count
                    </Col>

                    <Col>
                        {data.access_count}
                    </Col>
                </Row>

                <Timestamps timestamps={data.access_timestamps} />
            </>}
        </Container>
    )
}

export default Analytics