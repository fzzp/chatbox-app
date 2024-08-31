import React from 'react'
import {Modal, Button} from 'react-bootstrap';

type AddModalProps = {
    showModal: boolean;
    setShowModal: React.Dispatch<React.SetStateAction<boolean>>;
}

function AddModal({showModal, setShowModal}:AddModalProps) {
    const handleClose = () => setShowModal(false);
    
    return (
        <Modal
            show={showModal}
            onHide={()=>setShowModal(false)}
            backdrop="static"
            keyboard={false}
        >
            <Modal.Header closeButton>
                <Modal.Title>Modal title</Modal.Title>
            </Modal.Header>
            <Modal.Body>
               内容
            </Modal.Body>
            <Modal.Footer>
                <Button variant="secondary" onClick={handleClose}>
                    Close
                </Button>
                <Button variant="primary">Understood</Button>
            </Modal.Footer>
        </Modal>
    )
}

export default AddModal