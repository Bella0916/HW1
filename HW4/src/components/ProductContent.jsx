import React, {useEffect} from "react";
import BTable from 'react-bootstrap/Table';
import { useTable } from 'react-table'
import { useForm } from "react-hook-form";
import Form from "react-bootstrap/Form";
import Button from "react-bootstrap/Button";
import Stack from 'react-bootstrap/Stack';
import { InputGroup } from "react-bootstrap";



function Table({ columns, data }) {
    // Use the state and functions returned from useTable to build your UI
    const {
      getTableProps,
      getTableBodyProps,
      headerGroups,
      rows,
      prepareRow,
    } = useTable({
      columns,
      data,
    })
  
    // Render the UI for your table
    return (
      <BTable striped bordered hover size="sm"  {...getTableProps()}>
        <thead>
          {headerGroups.map(headerGroup => (
            <tr {...headerGroup.getHeaderGroupProps()}>
              {headerGroup.headers.map(column => (
                <th {...column.getHeaderProps()}>{column.render('Header')}</th>
              ))}
            </tr>
          ))}
        </thead>
        <tbody {...getTableBodyProps()}>
          {rows.map((row, i) => {
            prepareRow(row)
            return (
              <tr {...row.getRowProps()}>
                {row.cells.map(cell => {
                  return <td {...cell.getCellProps()}>{cell.render('Cell')}</td>
                })}
              </tr>
            )
          })}
        </tbody>
      </BTable>
    )
  }

export const ProductContent = (props) => {
    const [product, setProduct] = React.useState(null);
    const [data, setData] = React.useState([]);
    const [selectDBSource, setSelectDBSource] = React.useState(null);
    const { register, handleSubmit} = useForm();
    
    const updateProductSelection = async (data) => {
      const selectedProductId = data.productId; 
      let resJson = await getProductsById(selectedProductId);
      console.log(resJson);
      setData([resJson]);
      if(resJson.length == 0){
          alert("查無資料");
      }
  }

    const generateFormComps = (obj) => {
      if (!obj) return null;
      const { key, option, label } = obj;

      return (
          <InputGroup key={key}>
              <InputGroup.Text id={key}>{label}</InputGroup.Text>
              <Form.Select className="w-25" {...register(key)}>
                  {option.map(({ id, value }) => <option value={value} key={value}>{id}</option>)}
              </Form.Select>
          </InputGroup>
      );

  };

    useEffect(() => {
        setProduct(props.product);
    }, [props.product]);
    
    useEffect(() => {
        updateProductsWithTable();
        updateProductsOption();
    }, []);

    async function updateProductsOption() {
      const resJson = await getProducts();
      resJson.forEach(element => {
          console.log(element);
      });
      console.log(resJson);
      let options = [];
      resJson.forEach(element => {
        options.push({id:element.id, value:element.id});
      });
      let productsOption = {
        key:"productId", 
        option:options, 
        label:"產品"
      };
      console.log(productsOption);
      setSelectDBSource(productsOption);

  }

    function getProducts(){
        const headers = new Headers()
        const options = {
            method: "GET",
            headers: headers,
            //mode: 'no-cors'
        };
        const uri = 'http://localhost:8080/products'
        
        return fetch(uri, options).then(response => response.json())
    }

    function getProductsById(productId){
      const headers = new Headers()
      const options = {
          method: "GET",
          headers: headers,
          //mode: 'no-cors'
      };
      const uri = 'http://localhost:8080/products/'+ productId
      
      return fetch(uri, options).then(response => response.json())
  }

    async function updateProductsWithTable(){
        let resJson = await getProducts();
        resJson.forEach(element => {
            console.log(element);
        });
        console.log(resJson);
        setData(resJson);
    }

    const columns = React.useMemo(
        () => [
          {
            Header: 'Product ID',
            columns: [
              {
                Header: 'Product ID',
                accessor: 'id',
              },
            ],
          },
          {
            Header: 'Name',
            columns: [
              {
                Header: 'Name',
                accessor:'name',
              },
            ],
          },
          {
            Header: 'Price',
            columns: [
              {
                Header: 'Price',
                accessor:'price',
              },
              
            ],
          }
        ],
        []
      )
    
    
    return (
      <div id="ProductContent-div">
        <p style={{ textAlign: 'center' }}>
          ProductContent
        </p>
        {product}
        <div>
        <Form onSubmit={handleSubmit(updateProductSelection)}>
                <Stack direction="horizontal" gap={3}>
                    {generateFormComps(selectDBSource)}
                    <Button type="submit" className="btn btn-dark">查詢</Button>
                </Stack>
            </Form>
        </div>
        <div>
          <Table columns={columns} data={data} />
        </div>
      </div>
    
    );
}
    
