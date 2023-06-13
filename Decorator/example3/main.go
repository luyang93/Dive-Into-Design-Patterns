package main

import "fmt"

type DataSource interface {
	WriteData(data string)
	ReadData() string
}

type FileDataSource struct {
	filename string
	data     string
}

func NewFileDataSource(filename string) *FileDataSource {
	return &FileDataSource{filename: filename}
}

func (f *FileDataSource) WriteData(data string) {
	fmt.Println("Writing data to", f.filename)
	f.data = data
}

func (f *FileDataSource) ReadData() string {
	fmt.Println("Reading data from", f.filename)
	return f.data
}

type DataSourceDecorator struct {
	wrappee DataSource
}

func NewDataSourceDecorator(source DataSource) *DataSourceDecorator {
	return &DataSourceDecorator{wrappee: source}
}

func (d *DataSourceDecorator) WriteData(data string) {
	d.wrappee.WriteData(data)
}

func (d *DataSourceDecorator) ReadData() string {
	return d.wrappee.ReadData()
}

type EncryptionDecorator struct {
	*DataSourceDecorator
}

func NewEncryptionDecorator(source DataSource) *EncryptionDecorator {
	return &EncryptionDecorator{NewDataSourceDecorator(source)}
}

func (ed *EncryptionDecorator) WriteData(data string) {
	fmt.Println("encrypted")
	ed.DataSourceDecorator.WriteData(data)
}

func (ed *EncryptionDecorator) ReadData() string {
	fmt.Println("decrypted")
	return ed.DataSourceDecorator.ReadData()
}

type CompressionDecorator struct {
	*DataSourceDecorator
}

func NewCompressionDecorator(source DataSource) *CompressionDecorator {
	return &CompressionDecorator{NewDataSourceDecorator(source)}
}

func (cd *CompressionDecorator) WriteData(data string) {
	fmt.Println("compressed")
	cd.DataSourceDecorator.WriteData(data)
}

func (cd *CompressionDecorator) ReadData() string {
	fmt.Println("decompressed")
	return cd.DataSourceDecorator.ReadData()
}

func main() {
	source := NewFileDataSource("file.dat")
	source.WriteData("salaryRecords")
	fmt.Println(source.ReadData())
	fmt.Println("=======")

	source1 := NewDataSourceDecorator(source)
	source1.WriteData("salaryRecords")
	fmt.Println(source1.ReadData())
	fmt.Println("=======")

	source2 := NewEncryptionDecorator(source)
	source2.WriteData("salaryRecords")
	fmt.Println(source2.ReadData())
	fmt.Println("=======")

	source3 := NewCompressionDecorator(source)
	source3.WriteData("salaryRecords")
	fmt.Println(source3.ReadData())
	fmt.Println("=======")

	source4 := NewCompressionDecorator(source2)
	source4.WriteData("salaryRecords")
	fmt.Println(source4.ReadData())
	fmt.Println("=======")

	source5 := NewEncryptionDecorator(source3)
	source5.WriteData("salaryRecords")
	fmt.Println(source5.ReadData())
	fmt.Println("=======")
}
