import { Button, Tooltip } from "@mui/material";
import { getGridStringOperators, GridColDef, GridRenderCellParams } from "@mui/x-data-grid";

import DeleteIcon from '@mui/icons-material/Delete';
import EditOutlinedIcon from '@mui/icons-material/EditOutlined';
import moment from "moment";
import { NavigateFunction } from 'react-router-dom';

type DeleteFunc = (params: GridRenderCellParams) => void;

type ListColumnsType = {
  navigate: NavigateFunction
  onDelete?: DeleteFunc
}

export const ListColumns = ({
  navigate,
  onDelete,
}: ListColumnsType): GridColDef[] => [
    {
      field: 'id',
      headerName: 'Name',
      width: 400,
      sortable: true,
      filterable: true,
      filterOperators: getGridStringOperators().filter(
        (operator) => operator.value === 'contains',
      ),
    },
    {
      field: 'created_at',
      headerName: 'Creation date',
      width: 150,
      sortable: true,
      filterable: false,
      renderCell: (params: GridRenderCellParams) => {
        if (params.row.created_at.startsWith('0001-01-01')) {
          return (<i>Unknown</i>);
        }

        const date = moment(params.row.created_at);
        return (
          <div title={`${date.format('L')} at ${date.format('LT')}`}>
            {date.fromNow()}
          </div>
        )
      },
    },
    {
      field: 'updated_at',
      headerName: 'Update date',
      width: 150,
      sortable: true,
      filterable: false,
      renderCell: (params: GridRenderCellParams) => {
        if (params.row.updated_at.startsWith('0001-01-01')) {
          return (<i>Unknown</i>);
        }

        const date = moment(params.row.updated_at);
        return (
          <div title={`${date.format('L')} at ${date.format('LT')}`}>
            {date.fromNow()}
          </div>
        )
      },
    },
    {
      field: 'action',
      width: 250,
      type: 'actions',
      headerName: 'Actions',
      renderCell: (params: GridRenderCellParams) => (
        <>
          <Button
            variant='contained'
            size='small'
            color='primary'
            startIcon={(<EditOutlinedIcon />)}
            style={{ marginRight: 10 }}
            onClick={() => navigate('/principals/edit/' + params.row.id)}
          >
            Edit
          </Button>
          
          {params.row.is_locked ? (
            <Tooltip title='This principal is locked' placement='right'>
              <span>
                <Button
                  disabled
                  variant='text'
                  size='small'
                >
                  <DeleteIcon />
                </Button>
              </span>
            </Tooltip>          
          ) : (
            <Tooltip title='Delete' placement='right'>
              <Button
                variant='text'
                size='small'
                color='error'
                onClick={() => {
                  if (onDelete !== undefined) {
                    onDelete(params);
                  }
                }}
              >
                <DeleteIcon />
              </Button>
            </Tooltip>
          )}
        </>
      )
    },
];