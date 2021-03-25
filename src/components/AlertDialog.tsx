import React from 'react';
import Button from '@material-ui/core/Button';
import Dialog from '@material-ui/core/Dialog';
import DialogActions from '@material-ui/core/DialogActions';
import DialogContent from '@material-ui/core/DialogContent';
import DialogContentText from '@material-ui/core/DialogContentText';
import DialogTitle from '@material-ui/core/DialogTitle';

type AlertDialogProps = {
    isOpen: boolean,
    setIsOpen: Function,
    title: string,
    body: string,
    action: Function
  }

export default function AlertDialog({isOpen, setIsOpen, title, body, action}: AlertDialogProps) {
  
    const handleClose = () => {
        setIsOpen(false);
    };

    return (
        <Dialog
            open={isOpen}
            onClose={handleClose}
            aria-labelledby="alert-dialog-title"
            aria-describedby="alert-dialog-description"
        >
            <DialogTitle id="alert-dialog-title">{title}</DialogTitle>
            <DialogContent>
            <DialogContentText id="alert-dialog-description">
                {body}
            </DialogContentText>
            </DialogContent>
            <DialogActions>
            <Button onClick={handleClose} color="primary">
                No
            </Button>
            <Button onClick={()=>{action();handleClose();}} color="primary" autoFocus>
                Yes
            </Button>
            </DialogActions>
        </Dialog>
    );
}
