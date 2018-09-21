// +build linux,arm64
package pow

import (
	"github.com/iotaledger/giota/curl"
	"github.com/iotaledger/giota/transaction"
	"github.com/iotaledger/giota/trinary"
	"testing"
	"time"
)

func testPowCARM64(t *testing.T) float64 {
	var txTrytes trinary.Trytes = "SISEZJUUKSTSX9KVQGXSYYLNDIBJDVRZSOFEHWJSDZLNUUNBDLHUODEGFZQTKOEXUMMQTOREUWQCSGGWRKALQDDZCQN9LBIEVKBFDCWBIDWD9DGVOJVCNUNWDDZFCIOICZZF9KIAYDCSKJWE99UPPLUQPUSWTDKTSSTJAQNYATUTXZPA9CCJRRNIRWXTAR9ECVYXC9AOHXHYVOS9LWDUOH9SDUAQBEYTMJIMUHJTGUSQTFPRLLXIDKOVZMONJHXPCD9FYLW9PN9LLPQBJRSEKVKKJB9JRTZCXSDBMJYAKDX99EGNLFZPKIADJQEIMCKRFQKIHGCJAHPL9JFJF9PHRKPCHBPN9LYQSC9TXOXAI9WBDIBNGFPLQS9BHTEVROMCAXXAXPVBAP9URJXIVZXIWWCMVDXGAFZOIRTJIMNIZEPGFMWXWOWRDUMHFRKL9LV9VJQIRZPVJSSKHXHHVZLRZYHGWQAVL9BMWKKFGZQEYJNCGROYYDIDULQVSXGVLTTZRLPSKPVIURJ9CJBTNAYCPHQTWTTKHXPABTYYCCVAZATEVED9PBJQTNOQEQQBTSATZJTVUTZPUWDYKROBROUVSPMDLUMEZWMPESEMQPSVTDZKATUTOAEVWCW9HIKKHMOQYJOUYLTFPERSKBVWARHGJNKUWGFZYF9WSTEHEQWCA9DTOTOTNDFGAEABKKBKEFLDELEOYPZTCVKOBIWA9HWTCQT9IGYVFAFAOLOJMRDZKCBYOCPGEGGZL9CGFURM9FJBLGLZJILNSFOBXLQOZWVLAZUFLGQNCAVJTBGVLZETETWGXLPSPWMMAEGORSDGPUSFRQ9AVWWZCFNKSAHIKJOMEWCCFGVYSDYNIXYYTKJTOKZUGLKNEXHWQ9HVFVJUGJJEDQACTWPSFOONTNCJRDQBSCGXVKWZIGDK9RGHKAHSTOJDJEHIAOF9MFLAZJXLUGQUAUGKQGQIXXNLAPRQNTNVDGXVZBSEFXVRR9ZQIZEWPXZFMXLJFTFKEPPAFJTMBLBWYAWJEIHUNATL9EHIJQTCCMQFHILGHGEVXKHDCNMAHDPUGBQYYBF9CRIKDVZZ9KIFELUUKPXPRIFVTZPXRBKJBRLEGUJKXZPYGXRKOAHROFXENAUAYOSQBJGMMHIDUNSYYGQSDJDKMPNBPTUWMIYZCWABYLDMTXAGWFYEXRGLOYVPNSOVYITEPCXMTMPVLBQPBNQUBITEM99KVRTPNAAWPR9RQYBLFZDVWYDJXQRGTVAFVE99KE9YSCETBIELIWPKZYFARSPVLTDKEAKLCKULZHLKOQZMVLFLF9QHT9LLS9QQODSFYUIPKSBVSKAJMVW9QUILQSKHZMAXGVHUJBMTATPIDHJVUBZWUOYNOOMEJVOUXHACUHDVKZ9ZDTSIHQOTOVUMEISMA9VZIFQTPBXXDHDLVLKZZHLYLPIE9SKOEJXAFDKICOYIOVVAEXC9VZSFSDTSHVEOSHIT9JHMBBPQTRGOREIYQSBCMHJQIXTTQWOCKMCSGBRTJRRYWPXAGELIFPG9YX9FNNYGSJXJYTHIMWSXZH9JQIYXKFXEOHOE9YNHJIDAJUGPENZHOIFEHBSCQITVFHUOESVXOJPCNTUZR9LVQCXYUW9DITEXPG9KWYMBZQQCESNFVUOBQGCRRKFHOEKTHDHUNRXADXUMCWFJMZTMHN9VWLZATB9FF9HBGLFITNNVFCQICPRSGVFAATWYJT9GUJIAHNNJBECYSWSGEJYLHJPUOYESLVIELBMSLRZJLPKDKFGAJSSWZCQDLFDEXWAPILHLNHKCRMPLQUYESAEIWWNBCEIYSOHKPILTXPAFIZ9JMKFKJHTLHRHGZQLCEVJJMJHWTUKMKOWTZWGVZGQAOAKVGXZEZBMYPVWUGYJBIFXBACZLADFFBZIXKWSZLDOCGRQAZDCFPRAZYXUMNRJ9UKUKRAVSVMCENDJABZITDQLNCXZNXCOHKLATFFXKP9FFDYSAXISISMVYPXPWYPVEAYRNAITWJSTGXRAMMZIZF9IUORREWSFUNZOXDVCMBZJAET9PVHCQTMDTVVXLXDIXFSHPXWKBZBDJAAXSDEFXPARBU9GJJABPMCD9LGQJLRIYKGQORGCDDABAIAQC9MZDQLXFSAOLNYMWCJODEEUSIHEVHQPAIFQL9ECBBVZPHYU9HDBOYXTKWOIRGHUJMVV9UKHHREDIU9CRZFUZKAMUVRIEMKEKIMAGXSMGTEJWCWWAMRPWNINTETOTRMODTORVEURRY9RTDYQIEW99999999999999999999999999999999999999999999CMRKHWD99A99999999C99999999TNFAKVBFHHMKQKKSNJRLDIYUIGOMEOADJLNS9JGKGUIHZHIUDNQMVYCA9SZCLQOEVJPUGQGWTMETLGMUQMAKHHHHTBHVWYSJSXRVBRMHVV9WUTNMNFVDWLHQGFELTKZOISREPUJXNRBIAQVQWCCKB9DEZEXS999999M9EZGRXJ9WYSZXNDZBAJZMJ9VAMUWWWANGIVFKCUNRB9GLZZKRIMEFUK9KEFZXYDGBQJIU9SQUM999999999999999999999999999999999999999999999999999999999999999999999999999999999999999"

	s := time.Now()
	nonce, err := PowCARM64(txTrytes, 14)
	ti := time.Now().Sub(s)
	if err != nil {
		t.Fatal(err)
	}

	txTrytes = txTrytes[:len(txTrytes)-transaction.NonceTrinarySize/3] + nonce
	h := curl.Hash(txTrytes)
	if h[len(h)-4:] != "9999" {
		t.Error("pow is illegal", h)
	}

	return float64(countCARM64) / 1000 / ti.Seconds()
}

func TestPowCARM64(t *testing.T) {
	_proc := PowProcs

	tests := []struct {
		name     string
		powProcs int
	}{
		{
			name:     "test plain PowCARM64 without setting PowProcs",
			powProcs: PowProcs,
		},
		{
			name:     "test with PowProcs = 1",
			powProcs: 1,
		},
		{
			name:     "test with PowProcs = 32",
			powProcs: 32,
		},
		{
			name:     "test with PowProcs = 64",
			powProcs: 64,
		},
	}

	for _, tt := range tests {
		PowProcs = tt.powProcs
		sp := testPowCARM64(t)
		t.Logf("%s: %d kH/sec on SEE PoW", tt.name, int(sp))
	}

	PowProcs = _proc
}
