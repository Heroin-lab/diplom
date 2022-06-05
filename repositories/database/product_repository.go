package database

import (
	logger "github.com/Heroin-lab/heroin-logger/v3"
	"github.com/Heroin-lab/nixProject/models"
)

type ProductRepose struct {
	storage *Storage
}

func (r *ProductRepose) GetCrimesById(userId int) ([]*models.Crimes, error) {
	u := &models.Crimes{}

	allProdSql, err := r.storage.DB.Query("SELECT lat, lon, title_criminal, crime_number FROM offenses INNER JOIN criminal_code cc on offenses.crime_code_id = cc.id_criminal WHERE user_id= ?", userId)
	if err != nil {
		return nil, err
	}
	defer allProdSql.Close()

	allСrimes := make([]*models.Crimes, 0)

	for allProdSql.Next() {

		err = allProdSql.Scan(
			&u.Lat,
			&u.Lon,
			&u.TitleCriminal,
			&u.CrimeNumber,
		)
		if err != nil {
			return nil, err
		}

		allСrimes = append(allСrimes, &models.Crimes{
			Lat:           u.Lat,
			Lon:           u.Lon,
			TitleCriminal: u.TitleCriminal,
			CrimeNumber:   u.CrimeNumber,
		})
	}

	return allСrimes, nil
}

func (r *ProductRepose) InsertItem(p *models.Products) (*models.Products, error) {
	_, err := r.storage.DB.Exec("INSERT INTO products (product_name, type_id, price, img, supplier_id)\n "+
		"VALUES (?, ?, ?, ?, ?)",
		p.Product_name,
		p.Type_id,
		p.Price,
		p.Img,
		p.Supplier_id)
	if err != nil {
		return nil, err
	}

	logger.Info("Row with name '" + p.Product_name + "' was successfully added to PRODUCT table!")
	return p, nil
}

func (r *ProductRepose) DeleteItem(stringToDelete string) error {
	_, err := r.storage.DB.Exec("DELETE FROM products WHERE id=?",
		stringToDelete,
	)
	if err != nil {
		return err
	}

	logger.Info("Row with name '" + stringToDelete + "' was successfully deleted from PRODUCTS table!")
	return nil
}

func (r *ProductRepose) UpdateItem(p *models.Products) error {
	rows, err := r.storage.DB.Query("UPDATE products\n"+
		"SET product_name=?, type_id=?,\n"+
		"price=?, img=?,\n"+
		"supplier_id=?\n"+
		"WHERE id=?",
		p.Product_name,
		p.Type_id,
		p.Price,
		p.Img,
		p.Supplier_id,
		p.Id)
	if err != nil {
		return err
	}
	defer rows.Close()

	logger.Info("Row in PRODUCTS table was updated! RowID=", p.Id)
	return nil
}
