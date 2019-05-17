package checker

import (
	"fmt"
	"time"

	"github.com/pkg/errors"
	iconf "github.com/wncbb/check_api_diff/internal/conf"
	istorage "github.com/wncbb/check_api_diff/internal/storage"
)

func WriteDiffToFile(reqItem *iconf.ReqItem, content string, diffInfo *DiffInfo) error {
	if reqItem == nil {
		return errors.New("WriteDiffToFile's parameter reqItem should not be nil")
	}

	filePath := reqItem.Prefix
	fileName := reqItem.Name + ".diff"

	infoStr := fmt.Sprintf("# Info %s\n", time.Now())
	infoStr += fmt.Sprintf("# AddNum: %d, DelNum: %d\n", diffInfo.AddNum, diffInfo.DelNum)
	infoStr += fmt.Sprintf("# Path: %s\n", filePath)
	infoStr += fmt.Sprintf("# Name: %s\n", fileName)
	infoStr += fmt.Sprintf("# Description: %s \n", reqItem.Description)
	infoStr += "# Request:\n"
	infoStr += reqItem.Request.ShowComment("#   ")

	err := istorage.WriteFile(fileName, filePath, []byte(infoStr+content))
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
