package scarab

import (
	"context"
	"fmt"

	pbmodel "github.com/prometheus/client_model/go"
	"github.com/prometheus/prometheus/pkg/labels"
	"github.com/prometheus/prometheus/storage"
	"github.com/prometheus/prometheus/tsdb"
)

func (mw *metricsWriter) writeMetric(app storage.Appender, ls labels.Labels, v float64) {
	ts := int64(0)
	app.Append(0, ls, ts, v)
}

func (mw *metricsWriter) writeMetricFamily(app storage.Appender, mf *pbmodel.MetricFamily, ls labels.Labels) {
	lb := labels.NewBuilder(ls)
	lb.Set(labels.MetricName, *mf.Name)
	for i := range mf.Metric {
		switch *mf.Type {
		case pbmodel.MetricType_GAUGE:
			mw.writeMetric(app, lb.Labels(), *mf.Metric[i].Gauge.Value)
		case pbmodel.MetricType_COUNTER:
			mw.writeMetric(app, lb.Labels(), *mf.Metric[i].Counter.Value)
		case pbmodel.MetricType_UNTYPED:
			mw.writeMetric(app, lb.Labels(), *mf.Metric[i].Untyped.Value)
		case pbmodel.MetricType_HISTOGRAM:
			lb.Set(labels.MetricName, *mf.Name+"_count")
			mw.writeMetric(app, lb.Labels(), float64(*mf.Metric[i].Histogram.SampleCount))
			lb.Set(labels.MetricName, *mf.Name+"_sum")
			mw.writeMetric(app, lb.Labels(), float64(*mf.Metric[i].Histogram.SampleSum))
			for j := range mf.Metric[i].Histogram.Bucket {
				lb.Reset(ls)
				lb.Set(labels.MetricName, *mf.Name+"_bucket")
				lb.Set(labels.BucketLabel, fmt.Sprintf("%v", *mf.Metric[i].Histogram.Bucket[j].UpperBound))
				mw.writeMetric(app, lb.Labels(), float64(*mf.Metric[i].Histogram.Bucket[j].CumulativeCount))
			}
		case pbmodel.MetricType_SUMMARY:
			lb.Set(labels.MetricName, *mf.Name+"_count")
			mw.writeMetric(app, lb.Labels(), float64(*mf.Metric[i].Summary.SampleCount))
			lb.Set(labels.MetricName, *mf.Name+"_sum")
			mw.writeMetric(app, lb.Labels(), float64(*mf.Metric[i].Summary.SampleSum))
			for j := range mf.Metric[i].Summary.Quantile {
				lb.Reset(ls)
				lb.Set(labels.MetricName, *mf.Name)
				lb.Set("quantile", fmt.Sprintf("%v", *mf.Metric[i].Summary.Quantile[j].Quantile))
				mw.writeMetric(app, lb.Labels(), float64(*mf.Metric[i].Summary.Quantile[j].Value))
			}
		default:
		}
	}

}

type metricsWriter struct {
	db *tsdb.DB
}

/*
	db, err := tsdb.Open("tsdb/", nil, nil, nil, nil)
	if err != nil {
		return nil, err
	}
*/

func (mw *metricsWriter) WriteMetrics(ctx context.Context, ms []*pbmodel.MetricFamily, ls labels.Labels) error {
	app := mw.db.Appender(ctx)
	for _, mf := range ms {
		mw.writeMetricFamily(app, mf, ls)
	}
	app.Commit()

	return nil
}
