import React, { useState, useEffect } from "react";
import { makeStyles, Theme, createStyles } from "@material-ui/core/styles";
import Typography from '@material-ui/core/Typography';
import Container from '@material-ui/core/Container';
import { Box, Paper } from "@material-ui/core";
import Divider from "@material-ui/core/Divider";
import Grid from '@material-ui/core/Grid';
import MenuItem from '@material-ui/core/MenuItem';
import MuiAlert, { AlertProps } from "@material-ui/lab/Alert";
import Select from '@material-ui/core/Select';
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';
import Snackbar from "@material-ui/core/Snackbar";
import { Link as RouterLink } from "react-router-dom";
import { MedicalProductInterface } from "../models/MedicalProduct";
import { TreatmentRecordInterface } from "../models/TreatmentRecord";
import { UserInterface } from "../models/User";
import { MedRecordInterface } from "../models/MedRecord";


const useStyles = makeStyles((theme: Theme) =>

    createStyles({

        root: { flexGrow: 1 },

        container: { marginTop: theme.spacing(2) },

        paper: { padding: theme.spacing(2), color: theme.palette.text.secondary },

        table: { minWidth: 20 },

        textField: {
            marginLeft: theme.spacing(1),
            marginRight: theme.spacing(1),
            width: 200,
        },
        formControl: {
            margin: theme.spacing(1),
            minWidth: 120,
        },
        selectEmpty: {
            marginTop: theme.spacing(2),
        },

    })

);


const Alert = (props: AlertProps) => {
    return <MuiAlert elevation={6} variant="filled" {...props} />;
};


export default function MecRec() {
    const classes = useStyles();

    const [MedRecord, setMedRecord] = useState<Partial <MedRecordInterface>>({});
    
    const handleChange = (event: React.ChangeEvent<{ name?: string; value: unknown }>
    ) => {
        const name = event.target.name as keyof typeof MedRecord;
        setMedRecord({
          ...MedRecord,
          [name]: event.target.value,
        });
        if (event.target.name === "TreatmentRecordID") {
            setSelectedTreatment(treatmentRecord.find(tr => tr.ID === event.target.value));
        }
    }

    const [medicalProduct, setmedicalProduct] = useState<MedicalProductInterface[]>([]);
    const getmedicalProduct = async () => {
        const apiUrl = "http://localhost:8080/api/MedicalProduct/";
        const requestOptions = {
            method: "GET",
            headers: { "Content-Type": "application/json" },
        }

        fetch(apiUrl, requestOptions)
            .then((response) => response.json())
            .then((res) => {
                console.log(res.data);
                if (res.data) {
                    setmedicalProduct(res.data)
                } else {
                    console.log("else")
                }
            });
    }
    
    const [treatmentRecord, settreatmentRecord] = useState<TreatmentRecordInterface[]>([]);
    const gettreatmentRecord = async () => {
        const apiUrl = "http://localhost:8080/api/TreatmentRecord/";
        const requestOptions = {
            method: "GET",
            headers: { "Content-Type": "application/json" },
        }

        fetch(apiUrl, requestOptions)
            .then((response) => response.json())
            .then((res) => {
                console.log(res.data);
                if (res.data) {
                    settreatmentRecord(res.data)
                } else {
                    console.log("else")
                }
            });
    }

    const [user, setuser] = useState<UserInterface[]>([]);
    const getuser = async () => {
        const apiUrl = "http://localhost:8080/api/User/";
        const requestOptions = {
            method: "GET",
            headers: { "Content-Type": "application/json" },
        }

        fetch(apiUrl, requestOptions)
            .then((response) => response.json())
            .then((res) => {
                console.log(res.data);
                if (res.data) {
                    setuser(res.data)
                } else {
                    console.log("else")
                }
            });
    }

    // ตัวแปร selectedTreatment เก็บ treatmentRecord จากการค้นหาด้วย MedRecord.treatmentRecordID
    const [selectedTreatment, setSelectedTreatment] = React.useState<TreatmentRecordInterface>();


    const [success, setSuccess] = useState(false);
    const [error, setError] = useState(false);
    const handleClose = (event?: React.SyntheticEvent, reason?: string) => {
        if (reason === "clickaway") {
            return;
        }
        setSuccess(false);
        setError(false);
    };

    function submit() {
        let data = {
            Amount: typeof MedRecord.Amount === "string" ? parseInt(MedRecord.Amount) : 0,
            TreatmentRecordID: MedRecord.TreatmentRecordID,
            UserID: MedRecord.UserID,
            MedicalProductID: MedRecord.MedicalProductID,
        }

        const apiUrl = "http://localhost:8080/api/submit";
        const requestOptionsPost = {
            method: "POST",
            headers: {"Content-Type": "application/json"},
            body: JSON.stringify(data),
        }

        fetch(apiUrl, requestOptionsPost)
          .then((response) => response.json())
          .then((res) => {
            if (res.data) {
              setSuccess(true);
            } else {
              setError(true);
            }
        });

    }

    /* --- DEBUG --- */
    console.log("data_submit", MedRecord);
        
    useEffect(() => {
        getmedicalProduct();
        gettreatmentRecord();
        getuser();
    }, []);
    
    return (


        <Container className={classes.container} maxWidth="md">
            <Snackbar open={success} autoHideDuration={5000} onClose={handleClose}>
                <Alert onClose={handleClose} severity="success">
                    บันทึกข้อมูลสำเร็จ
                </Alert>
            </Snackbar>
            <Snackbar open={error} autoHideDuration={5000} onClose={handleClose}>
                <Alert onClose={handleClose} severity="error">
                    บันทึกข้อมูลไม่สำเร็จ
                </Alert>
            </Snackbar>
            
            <Paper className={classes.paper}>
                <Box display="flex">
                    <Box flexGrow={1}>
                        <Typography
                            component="h2"

                            variant="h6"

                            color="primary"

                            gutterBottom
                        >
                            ข้อมูลการบันทึกการจ่ายยาและเวชภัณฑ์
                        </Typography>

                    </Box>
                </Box>
                <Divider />


                <Grid container spacing={3} className={classes.root}>

                    <Grid item xs={12}>
                        <p>ชื่อผู้จ่ายยาและเวชภัณฑ์</p>
                        <Select
                            style = {{width:400}}
                            variant = "outlined"
                            defaultValue = {0}
                            value={MedRecord.UserID}
                            onChange={handleChange}
                            inputProps={{name: "UserID",}}
                        >
                            <MenuItem value={0} key={0}>เลือกผู้จ่ายยา</MenuItem>
                            {user.map((item: UserInterface) => (
                                <MenuItem value={item.ID} key={item.ID}>
                                    {item.Name}
                                </MenuItem>
                            ))}
                        </Select>
                    </Grid>


                    <Grid item xs={12}>
                        <p>เลขใบวินิฉัย</p>
                        <Select
                            style = {{width:400}}
                            variant = "outlined"
                            defaultValue = {0}
                            value={MedRecord.TreatmentRecordID}
                            onChange={handleChange}
                            inputProps={{name: "TreatmentRecordID"}}
                        >
                            <MenuItem value={0} key={0}>เลือกเลขใบวินิฉัย</MenuItem>
                            {treatmentRecord.map((item: TreatmentRecordInterface) => (
                                <MenuItem value={item.ID} key={item.ID}>
                                    {item.ID}
                                </MenuItem>
                            ))}
                        </Select>
                    </Grid>



                    <Grid item xs={6}>
                        <p>ชื่อ</p>
                        <TextField fullWidth variant="outlined" disabled value={selectedTreatment?.ScreeningRecord.Patient.Firstname} />
                    </Grid>
                    <Grid item xs={6}>
                        <p>นามสกุล</p>
                        <TextField fullWidth variant="outlined" disabled value={selectedTreatment?.ScreeningRecord.Patient.Lastname} />
                    </Grid>


                    <Grid item xs={6}>
                        <p>ชื่อยา</p>
                        <Select
                            fullWidth
                            variant = "outlined"
                            defaultValue = {0}
                            value={MedRecord.MedicalProductID}
                            onChange={handleChange}
                            inputProps={{name: "MedicalProductID"}}
                        >
                            <MenuItem value={0} key={0}>เลือกรายการยาและเวชภัณฑ์</MenuItem>
                            {medicalProduct.map((item: MedicalProductInterface) => (
                                <MenuItem value={item.ID} key={item.ID}>
                                    {item.Name}
                                </MenuItem>
                            ))}
                        </Select>
                    </Grid>


                    <Grid item xs={6}>
                        <p>จำนวน</p>
                        <TextField 
                            fullWidth 
                            id="outlined-basic" 
                            label="" 
                            variant="outlined" 
                            value={MedRecord.Amount} 
                            onChange={handleChange}
                            inputProps={{name: "Amount"}}
                        />
                    </Grid>

                    <Grid item xs={12}>

                        <Button component={RouterLink} to="/" variant="contained">

                            กลับ

                        </Button>

                        <Button

                            style={{ float: "right" }}

                            onClick={submit}

                            variant="contained"

                            color="primary"

                        >

                            บันทึก

                        </Button>

                    </Grid>



                </Grid>


            </Paper>

        </Container>

    )
}
